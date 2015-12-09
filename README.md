# Mandrill [![Build Status](https://travis-ci.org/jimtsao/mandrill.svg)](https://travis-ci.org/jimtsao/mandrill)
Mandrill is a Go(lang) API for Mandrill Email Service

# Import
    import "github.com/jimtsao/mandrill"

# Usage
    m := mandrill.NewMandrill("your-api-key")
    resp, err := m.Messages().SimpleSend("from@email.com", "to@email.com", "subject", "body")
    if err != nil {
        // handle error
    }