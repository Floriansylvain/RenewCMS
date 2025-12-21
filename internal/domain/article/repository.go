package article

type Repository interface {
	Get(id uint32) (Article, error)
	GetByName(name string) (Article, error)
	GetAll() []Article
	Create(post Article) (Article, error)
	UpdateBody(id uint32, body string) (Article, error)
	UpdateIsOnline(id uint32, isOnline bool) (Article, error)
	Delete(id uint32) error
	AddImage(postId uint32, imageId uint32) error
}
