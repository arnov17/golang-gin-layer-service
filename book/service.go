package book

type Service interface {
	FindAll() ([]Book, error)
	FindById(ID int) (Book error)
	Create(BookRequest BookRequest) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	book, err := s.repository.FindAll()
	return book, err
	// return s.repository.FindAll()
}

func (s *service) FindById(ID int) (Book, error) {
	var book Book
	err := s.repository.FindById(ID)
	return book, err
}

func (s *service) Create(BookRequest BookRequest) (Book, error) {
	price, _ := BookRequest.Price.Int64()

	book := Book{
		Title: BookRequest.Title,
		Price: int(price),
	}
	newBook, err := s.repository.Create(book)
	return newBook, err
}
