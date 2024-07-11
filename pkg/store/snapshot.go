package store

import (
	"encoding/json"
	"errors"
	"job-seek/pkg/request"
	seekGql "job-seek/pkg/request/seek_gql"
	"log"

	"github.com/cockroachdb/pebble"
)

var (
	db *pebble.DB

	ErrDBNotInit = errors.New("DB not init")
)

func keyUpperBound(b []byte) []byte {
	end := make([]byte, len(b))
	copy(end, b)
	for i := len(end) - 1; i >= 0; i-- {
		end[i] = end[i] + 1
		if end[i] != 0 {
			return end[:i+1]
		}
	}
	return nil // no upper-bound
}

func prefixIterOptions(prefix []byte) *pebble.IterOptions {
	return &pebble.IterOptions{
		LowerBound: prefix,
		UpperBound: keyUpperBound(prefix),
	}
}

func InitDB() {
	var err error
	db, err = pebble.Open("store", &pebble.Options{})
	if err != nil {
		log.Fatal(err)
	}
}

// / PostDetail
func CheckKeyPostDetailCache(postId string) (bool, error) {
	if db == nil {
		return false, nil
	}
	value, closer, err := db.Get([]byte("post-" + postId))
	if err != nil {
		return false, err
	}
	defer closer.Close()

	return value != nil, nil
}

func GetPostDetailCache(postId string) (*request.SeekPostDetails, error) {
	if db == nil {
		return nil, ErrDBNotInit
	}

	postDetail := new(request.SeekPostDetails)
	// b := new(bytes.Buffer)

	value, closer, err := db.Get([]byte("post-" + postId))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	// if err = ikea.Unpack(bytes.NewReader(value), postDetail); err != nil {
	if err = json.Unmarshal(value, postDetail); err != nil {
		return nil, err
	}

	return postDetail, nil
}

func LoadPostDetailCache() (*[]request.SeekPostDetails, error) {
	if db == nil {
		return nil, ErrDBNotInit
	}

	postDetail := []request.SeekPostDetails{}
	// b := new(bytes.Buffer)

	iter, _ := db.NewIter(prefixIterOptions([]byte("post-")))
	for iter.First(); iter.Valid(); iter.Next() {
		post := new(request.SeekPostDetails)
		if err := json.Unmarshal(iter.Value(), post); err != nil {
			log.Fatal(err)
		}
		postDetail = append(postDetail, *post)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &postDetail, nil
}

func SetPostDetailCache(postDetail *request.SeekPostDetails) error {
	if db == nil {
		return ErrDBNotInit
	}

	// b := new(bytes.Buffer)
	// ikea.Pack(b, postDetail)

	data, err := json.Marshal(postDetail)
	if err != nil {
		return err
	}

	err = db.Set([]byte("post-"+postDetail.PostId), data, pebble.Sync)

	if err != nil {
		return err
	}
	return nil
}

// /
func CheckKeyPostGQLCache(jobId string) (bool, error) {
	if db == nil {
		return false, nil
	}
	value, closer, err := db.Get([]byte("job-" + jobId))
	if err != nil {
		return false, err
	}
	defer closer.Close()

	return value != nil, nil
}

func GetPostGQLCache(jobId string) (*seekGql.Job, error) {
	if db == nil {
		return nil, ErrDBNotInit
	}

	jobDetail := new(seekGql.Job)

	value, closer, err := db.Get([]byte("job-" + jobId))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	// if err = ikea.Unpack(bytes.NewReader(value), jobDetail); err != nil {
	if err = json.Unmarshal(value, jobDetail); err != nil {
		return nil, err
	}

	return jobDetail, nil
}
func SetPostGQLCache(jobDetail *seekGql.Job) error {
	if db == nil {
		return ErrDBNotInit
	}

	// b := new(bytes.Buffer)
	// ikea.Pack(b, jobDetail)

	data, err := json.Marshal(jobDetail)
	if err != nil {
		return err
	}

	err = db.Set([]byte("job-"+jobDetail.ID), data, pebble.Sync)
	if err != nil {
		return err
	}
	return nil
}

// CompanyDetail
func CheckKeyCompanyDetailCache(companyId string) (bool, error) {
	if db == nil {
		return false, nil
	}

	value, closer, err := db.Get([]byte("c-" + companyId))
	if err != nil {
		return false, err
	}
	defer closer.Close()

	return value != nil, nil
}

func GetCompanyDetailCache(companyId string) (*request.SeekCompanyDetails, error) {
	if db == nil {
		return nil, ErrDBNotInit
	}

	companyDetail := new(request.SeekCompanyDetails)
	// b := new(bytes.Buffer)

	value, closer, err := db.Get([]byte("c-" + companyId))
	if err != nil {
		return nil, err
	}
	defer closer.Close()

	// if err = ikea.Unpack(bytes.NewReader(value), companyDetail); err != nil {
	if err = json.Unmarshal(value, companyDetail); err != nil {
		return nil, err
	}

	return companyDetail, nil
}

func LoadCompanyDetailCache() (*[]request.SeekCompanyDetails, error) {
	if db == nil {
		return nil, ErrDBNotInit
	}

	postDetail := []request.SeekCompanyDetails{}
	// b := new(bytes.Buffer)

	iter, _ := db.NewIter(prefixIterOptions([]byte("c-")))
	for iter.First(); iter.Valid(); iter.Next() {
		post := new(request.SeekCompanyDetails)
		if err := json.Unmarshal(iter.Value(), post); err != nil {
			log.Fatal(err)
		}
		postDetail = append(postDetail, *post)
	}
	if err := iter.Close(); err != nil {
		log.Fatal(err)
		return nil, err
	}

	return &postDetail, nil
}

func SetCompanyDetailCache(companyDetail *request.SeekCompanyDetails) error {
	if db == nil {
		return ErrDBNotInit
	}

	// b := new(bytes.Buffer)
	// ikea.Pack(b, companyDetail)

	data, err := json.Marshal(companyDetail)
	if err != nil {
		return err
	}

	err = db.Set([]byte("c-"+companyDetail.ReferenceId), data, pebble.Sync)

	if err != nil {
		return err
	}
	return nil
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}

func StorePoint() error {
	if db == nil {
		return ErrDBNotInit
	}
	db.Flush()

	return nil
}
