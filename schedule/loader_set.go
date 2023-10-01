package schedule

type LoaderSet struct {
	Loaders []Loader
}

func (l LoaderSet) List() ([]*Schedule, error) {
	result := make([]*Schedule, 0)
	for _, loader := range l.Loaders {
		loaderList, err := loader.List()
		if err != nil {
			return nil, err
		}
		for _, sch := range loaderList {
			result = append(result, sch)
		}
	}

	return result, nil
}
