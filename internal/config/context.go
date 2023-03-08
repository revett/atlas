package config

// ContextKey is a dedicated type for accessing context values via a key, to avoid collisions.
// https://www.calhoun.io/pitfalls-of-context-values-and-how-to-avoid-or-mitigate-them
type ContextKey string

// ContextConfigKey is the key used to store the Config in a context.Context.
const ContextConfigKey ContextKey = "atlas.config"
