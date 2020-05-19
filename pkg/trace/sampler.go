package trace

import (
	"github.com/uber/jaeger-client-go/config"
)

// ConstantSampler and other sampler types
const (
	ConstantSampler       = "const"
	ProbablilisticSampler = "probabilistic"
	RateLimitingSampler   = "ratelimiting"
	RepoteSampler         = "remote"
)

func NewSamplerConfig(sampler Sampler) config.SamplerConfig {
	switch sampler.Type {
	case ConstantSampler:
		return newConstantSamplerConfig(sampler.Constant)
	case ProbablilisticSampler:
		return newProbablisiticSamplerConfig(sampler.Probabilistic)
	case RateLimitingSampler:
		return newRateLimitingSamplerConfig(sampler.RateLimiting)
	case RepoteSampler:
		return newRemoteSamplerConfig(sampler.Remote)
	case "default":
		return config.SamplerConfig{}
	}

	return config.SamplerConfig{}
}

func newConstantSamplerConfig(constant Constant) config.SamplerConfig {
	sample := 0
	if constant.Enabled {
		sample = 1
	}

	return config.SamplerConfig{
		Type:  ConstantSampler,
		Param: float64(sample), // valid 0 - none / 1 - all
	}
}

func newProbablisiticSamplerConfig(prob Probablilistic) config.SamplerConfig {
	return config.SamplerConfig{
		Type:  ProbablilisticSampler,
		Param: prob.Probability, // float between 0-1.0
	}
}

func newRateLimitingSamplerConfig(rl RateLimiting) config.SamplerConfig {
	return config.SamplerConfig{
		Type:  RateLimitingSampler,
		Param: float64(rl.SpansPerSec), // valid 0 - none / 1 - all
	}
}

func newRemoteSamplerConfig(remote Remote) config.SamplerConfig {
	return config.SamplerConfig{
		Type:                     ProbablilisticSampler,
		Param:                    remote.Probability, // float between 0-1.0
		SamplingServerURL:        remote.ServerURL,
		SamplingRefreshInterval:  remote.RefreshInterval,
		MaxOperations:            remote.MaxOp,
		OperationNameLateBinding: remote.LateBinding,
	}
}
