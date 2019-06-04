package coolog

type Adapter string

type Config struct {
	adapter Adapter
	level   level
}
