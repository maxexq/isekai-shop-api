package web

type (
	Paginate struct {
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}

	PaginateOptional struct {
		Page int64 `query:"page" validate:"required,min=1"`
		Size int64 `query:"size" validate:"required,min=1,max=20"`
	}

	Pagination struct {
		Page      int64 `json:"page"`
		TotalPage int64 `json:"totalPage"`
	}
)
