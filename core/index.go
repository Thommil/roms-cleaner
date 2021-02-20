package core

import (
	"strings"

	"github.com/blevesearch/bleve/v2"
	"github.com/blevesearch/bleve/v2/mapping"
	"github.com/golang/glog"
)

// Index defines indexing API
type Index interface {
	Add(key string, data interface{}) error
	Search(query string) ([]SearchResult, error)
	Status() SearchStatus
	Close() error
}

// SearchStatus maintains current search result status
type SearchStatus struct {
	Total      int
	Failed     int
	Successful int
}

// SearchResult is the returned object by an Index.Search query
type SearchResult struct {
	Key   string
	Score float64
}

type bleveIndex struct {
	index  bleve.Index
	addInc int
	batch  *bleve.Batch
}

// CreateIndex instanciates a new Index implementation
func CreateIndex(excludedPaths []string) (Index, error) {
	glog.V(2).Infof("create new index")

	mapping := bleve.NewIndexMapping()
	if len(excludedPaths) > 0 {
		customMapping := bleve.NewDocumentMapping()
		for _, path := range excludedPaths {
			paths := strings.Split(path, ".")
			pathToMapping(paths, customMapping)
		}
		mapping.DefaultMapping = customMapping
	}
	index, err := bleve.NewMemOnly(mapping)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	batch := index.NewBatch()

	glog.V(2).Infof("index created")

	return &bleveIndex{
		index:  index,
		addInc: 0,
		batch:  batch,
	}, nil
}

func pathToMapping(path []string, documentMapping *mapping.DocumentMapping) error {
	localPath := strings.TrimSpace(path[0])
	if len(path) == 1 {
		ignoreMapping := bleve.NewTextFieldMapping()
		ignoreMapping.IncludeInAll = false
		ignoreMapping.Store = false
		documentMapping.AddFieldMappingsAt(localPath, ignoreMapping)
		documentMapping.AddSubDocumentMapping(localPath, mapping.NewDocumentDisabledMapping())
		return nil
	}

	if property, found := documentMapping.Properties[localPath]; !found {
		property = mapping.NewDocumentMapping()
		documentMapping.AddSubDocumentMapping(localPath, property)
		pathToMapping(path[1:], property)
	} else {
		pathToMapping(path[1:], property)
	}

	return nil
}

func (instance *bleveIndex) Add(key string, data interface{}) error {
	err := instance.batch.Index(key, data)

	if err != nil {
		glog.Error(err)
		return err
	}

	instance.addInc++

	if instance.addInc > 1000 {
		glog.V(2).Infof("flush batch")

		err = instance.index.Batch(instance.batch)
		instance.addInc = 0

		if err != nil {
			glog.Error(err)
			return err
		}
	}

	return nil
}

func (instance *bleveIndex) Search(query string) ([]SearchResult, error) {
	glog.V(2).Infof("new search : %s", query)

	if instance.addInc > 0 {
		err := instance.index.Batch(instance.batch)
		instance.addInc = 0

		if err != nil {
			glog.Error(err)
			return nil, err
		}
	}

	search := bleve.NewSearchRequest(bleve.NewQueryStringQuery(query))
	bleeveSearchResults, err := instance.index.Search(search)

	if err != nil {
		glog.Error(err)
		return nil, err
	}

	searchResults := make([]SearchResult, 0, len(bleeveSearchResults.Hits))

	for _, hit := range bleeveSearchResults.Hits {
		searchResults = append(searchResults, SearchResult{hit.ID, hit.Score})
	}

	glog.V(2).Infof("result : %v", searchResults)

	return searchResults, nil
}

func (instance *bleveIndex) Status() SearchStatus {
	return SearchStatus{
		Total:      instance.Status().Total,
		Successful: instance.Status().Successful,
		Failed:     instance.Status().Failed,
	}
}

func (instance *bleveIndex) Close() error {
	return instance.index.Close()
}
