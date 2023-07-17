package nostr

// SetExtra sets an out-of-the-spec value under the given key into the event object.
func (evt *Event) SetExtra(key string, value any) {
	if evt.extra == nil {
		evt.extra = make(map[string]any)
	}
	evt.extra[key] = value
}

// GetExtra tries to get a value under the given key that may be present in the event object
// but is hidden in the basic type since it is out of the spec.
func (evt Event) GetExtra(key string) any {
	// zero value of any is nil
	ival, _ := evt.extra[key]
	return ival
}

// GetExtraString is like [Event.GetExtra], but only works if the value is a string,
// otherwise returns the zero-value.
func (evt Event) GetExtraString(key string) (result string) {
	ival, ok := evt.extra[key]
	if !ok {
		return
	}
	val, ok := ival.(string)
	if !ok {
		return
	}
	return val
}

// GetExtraNumber is like [Event.GetExtra], but only works if the value is a float64,
// otherwise returns the zero-value.
func (evt Event) GetExtraNumber(key string) (result float64) {
	ival, ok := evt.extra[key]
	if !ok {
		return
	}

	switch val := ival.(type) {
	case float64:
		return val
	case int:
		return float64(val)
	case int64:
		return float64(val)
	}

	return
}

// GetExtraBoolean is like [Event.GetExtra], but only works if the value is a boolean,
// otherwise returns the zero-value.
func (evt Event) GetExtraBoolean(key string) (result bool) {
	ival, ok := evt.extra[key]
	if !ok {
		return
	}
	val, ok := ival.(bool)
	if !ok {
		return
	}
	return val
}
