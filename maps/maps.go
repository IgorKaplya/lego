package maps

type Dic map[string]string

type DicError string

func (e DicError) Error() string {
	return string(e)
}

const (
	ErrKeyNotFound      = DicError("key not found")
	ErrKeyAlreadyExists = DicError("key already exists")
)

func (d Dic) Search(key string) (value string, err error) {
	value, keyFound := d[key]
	if !keyFound {
		err = ErrKeyNotFound
	}
	return value, err
}

func (d *Dic) Add(key string, value string) error {
	_, keyFound := (*d)[key]
	if keyFound {
		return ErrKeyAlreadyExists
	}

	(*d)[key] = value

	return nil
}

func (d *Dic) Update(key string, value string) error {
	_, keyFound := (*d)[key]
	if !keyFound {
		return ErrKeyNotFound
	}

	(*d)[key] = value

	return nil
}

func (d *Dic) Delete(key string) error {
	_, keyFound := (*d)[key]

	if !keyFound {
		return ErrKeyNotFound
	}

	delete((*d), key)

	return nil
}
