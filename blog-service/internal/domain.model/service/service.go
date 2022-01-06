package service

import (
    "context"
    "sync"
)

/**
 * @author Rancho
 * @date 2021/12/10
 */

var (
    Service *DomainService
    once    sync.Once
)

type DomainService struct {
    *ExampleService
    *TagService
    *AuthService
}

type DomainServiceOption func(srv *DomainService)

func Init(ctx context.Context) {
    once.Do(func() {
        Service = NewDomainService(
            WithExampleService(ctx),
            WithTagService(ctx),
            WithAuthService(ctx),
        )
    })
}

func NewDomainService(opts ...DomainServiceOption) *DomainService {
    srv := &DomainService{}

    for _, opt := range opts {
        opt(srv)
    }
    return srv
}

func WithExampleService(ctx context.Context) DomainServiceOption {
    return func(s *DomainService) {
        s.ExampleService = NewExampleService(ctx)
    }
}

func WithTagService(ctx context.Context) DomainServiceOption {
    return func(s *DomainService) {
        s.TagService = NewTagService(ctx)
    }
}

func WithAuthService(ctx context.Context) DomainServiceOption {
    return func(s *DomainService) {
        s.AuthService = NewAuthService(ctx)
    }
}
