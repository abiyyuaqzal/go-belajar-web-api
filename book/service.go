package book

type Service interface {
	FindAll() ([]Book, error)
	FindByID(ID int) (Book, error)
	Create(bookRequest BookRequest) (Book, error)
	Delete(ID int) (Book, error)
	Update(book Book) (Book, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) FindAll() ([]Book, error) {
	// books, err := s.repository.FindAll()
	// return books, err

	return s.repository.FindAll()
}

func (s *service) FindByID(ID int) (Book, error) {
	return s.repository.FindByID(ID)
}

func (s *service) Create(bookRequest BookRequest) (Book, error) {
	price, _ := bookRequest.Price.Int64()
	rating, _ := bookRequest.Rating.Int64()

	book := Book{
		Title:       bookRequest.Title,
		Price:       int(price),
		Description: bookRequest.Description,
		Rating:      int(rating),
	}

	newBook, err := s.repository.Create(book)
	return newBook, err
}

func (s *service) Delete(ID int) (Book, error) {
	return s.repository.Delete(ID)
}

func (s *service) Update(book Book) (Book, error) {
	return s.repository.Update(book)
}
