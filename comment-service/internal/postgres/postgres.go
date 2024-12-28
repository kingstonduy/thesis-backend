package postgres

import (
	"context"
	"fmt"

	configuration "github.com/kingstonduy/comment-service/internal/bootstrap"
	"github.com/kingstonduy/comment-service/internal/domain"
	"github.com/kingstonduy/go-core/logger"
)

type productRepoImlp struct {
	db *configuration.PostgresCon
}

func NewProductRepoImpl(db *configuration.PostgresCon) domain.ICommentRepo {
	return &productRepoImlp{
		db: db,
	}
}

// AddComment implements domain.ICommentRepo.
func (p *productRepoImlp) AddComment(ctx context.Context, params domain.AddCommentParams) error {
	logger.Info(ctx, "AddComment start")
	defer logger.Info(ctx, "AddComment end")

	sqlQuery := `
        INSERT INTO public."COMMENT" (
            "PRODUCT_ID",
            "USER_ID",
            "REVIEW_EVALUATION",
            "REVIEW_DETAIL",
            "CREATED_AT",
            "UPDATED_AT"
        ) VALUES (
            $1,
            $2,
            $3,
            $4,
            CURRENT_TIMESTAMP,
            CURRENT_TIMESTAMP
        );
    `

	_, err := p.db.DB.Exec(ctx, sqlQuery,
		params.ProductID,
		params.UserID,
		params.ReviewEvaluation,
		params.ReviewDetail,
	)
	if err != nil {
		return fmt.Errorf("failed to add comment: %w", err)
	}

	return nil
}

// GetCommennt implements domain.ICommentRepo.
func (p *productRepoImlp) GetCommennt(ctx context.Context, params domain.GetCommentParamsIn) (domain.GetCommentParamsOut, error) {
	logger.Info(ctx, "GetCommennt start")
	defer logger.Info(ctx, "GetCommennt end")

	var sqlQuery string
	var queryArgs []interface{}

	// Construct the base query
	sqlQuery = `
        SELECT 
            c."COMMENT_ID",
            c."REVIEW_EVALUATION",
            c."REVIEW_DETAIL",
            c."CREATED_AT",
            c."USER_ID",
            u."USER_NAME",
            u."USER_IMAGE"
        FROM public."COMMENT" c
        INNER JOIN public."CUSTOMER" u ON c."USER_ID" = u."USER_ID"
        WHERE c."PRODUCT_ID" = $1
    `
	queryArgs = append(queryArgs, params.ProductID)

	// Add filtering logic
	if params.Filter != 0 {
		sqlQuery += ` AND c."REVIEW_EVALUATION" = $2`
		queryArgs = append(queryArgs, params.Filter)
	}

	// Add sorting logic
	if params.Sort == "asc" || params.Sort == "desc" {
		sqlQuery += ` ORDER BY c."CREATED_AT" ` + params.Sort
	} else {
		// Default sorting
		sqlQuery += ` ORDER BY c."CREATED_AT" DESC`
	}

	rows, err := p.db.DB.Query(ctx, sqlQuery, queryArgs...)
	if err != nil {
		return domain.GetCommentParamsOut{}, err
	}
	defer rows.Close()

	var comments []*domain.GetCommentParamsDetail
	for rows.Next() {
		var comment domain.GetCommentParamsDetail
		err := rows.Scan(
			&comment.CommentID,
			&comment.ReviewEvaluation,
			&comment.ReviewDetail,
			&comment.CreatedAt,
			&comment.UserID,
			&comment.UserName,
			&comment.UserImage,
		)
		if err != nil {
			return domain.GetCommentParamsOut{}, err
		}
		comments = append(comments, &comment)
	}

	if err = rows.Err(); err != nil {
		return domain.GetCommentParamsOut{}, err
	}

	return domain.GetCommentParamsOut{Comments: comments}, nil
}
