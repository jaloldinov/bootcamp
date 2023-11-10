package postgres

import (
	"context"
	"example-grpc-server/models"
	"fmt"

	"github.com/jackc/pgx/v4/pgxpool"
)

type biznesRepo struct {
	db *pgxpool.Pool
}

func NewBiznesRepo(db *pgxpool.Pool) *biznesRepo {
	return &biznesRepo{
		db: db,
	}
}

func (b *biznesRepo) GetTopStaff(c context.Context, req *models.TopStaffRequest) (*models.TopStaffResponse, error) {
	TopStaffs := models.TopStaffResponse{}
	TopStaffs.TopStaffs = make([]*models.TopStaff, 0)

	// query for CASHIER
	queryC := ` 
			SELECT  
   				 st.name, 
   				 b.name AS branch_name,
   				 SUM(s.price) AS total_sum,
   				 st.staff_type
			FROM sales AS s
			JOIN branches AS b ON b.id = s.branch_id
			JOIN staffs AS st ON st.id = s.cashier_id
			WHERE s.status = 'success'
		    AND date(s.created_at) >= $1 AND date(s.created_at) <= $2
				GROUP BY s.cashier_id, st.name, branch_name, st.staff_type
				ORDER BY total_sum LIMIT 1;
  `
	rowsC, err := b.db.Query(context.Background(), queryC, req.FromDate, req.ToDate)
	if err != nil {
		return nil, err
	}
	defer rowsC.Close()

	for rowsC.Next() {
		staff := models.TopStaff{}
		err := rowsC.Scan(
			&staff.Name,
			&staff.Branch,
			&staff.Total_Sum,
			&staff.StaffType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row cashier: %w", err)
		}
		TopStaffs.TopStaffs = append(TopStaffs.TopStaffs, &staff)
	}

	// query for SHOP_ASSISTANT
	queryA := ` 
		SELECT  
   			 st.name, 
   			 b.name AS branch_name,
   			 SUM(s.price) AS total_sum,
   			 st.staff_type
		FROM sales AS s
		JOIN branches AS b ON b.id = s.branch_id
		JOIN staffs AS st ON st.id = s.shop_assistant_id
		WHERE s.status = 'success'
 		    AND date(s.created_at) >= $1 AND date(s.created_at) <= $2
			GROUP BY s.cashier_id, st.name, branch_name, st.staff_type
			ORDER BY total_sum LIMIT 1;
  `

	rowsA, err := b.db.Query(context.Background(), queryA, req.FromDate, req.ToDate)
	if err != nil {
		return nil, err
	}
	defer rowsA.Close()

	for rowsA.Next() {
		staff := models.TopStaff{}
		err := rowsA.Scan(
			&staff.Name,
			&staff.Branch,
			&staff.Total_Sum,
			&staff.StaffType,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan row assistant: %w", err)
		}
		TopStaffs.TopStaffs = append(TopStaffs.TopStaffs, &staff)
	}
	return &TopStaffs, nil
}
