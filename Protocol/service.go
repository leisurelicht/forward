package Protocol

type Service interface {
    Run(args interface{})(err error)
}