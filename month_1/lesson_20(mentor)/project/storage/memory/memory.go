package memory

import "lesson_20/storage"

type store struct {
	branches    *branchRepo
	staffes     *staffRepo
	sales       *saleRepo
	transaction *transactionRepo
	staffTarifs *staffTarifRepo
}

func NewStorage(fileBranch, fileStaffes, fileSales, fileTransaction, fileTariffes string) storage.StorageI {
	return &store{
		branches:    NewBranchRepo(fileBranch),
		staffes:     NewStaffRepo(fileStaffes),
		sales:       NewSaleRepo(fileSales),
		transaction: NewTransactionRepo(fileTransaction),
		staffTarifs: NewStaffTarifRepo(fileTariffes),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffesI {
	return s.staffes
}

func (s *store) Sales() storage.SalesI {
	return s.sales
}

func (s *store) Transaction() storage.TransactionI {
	return s.transaction
}

func (s *store) StaffTarif() storage.StaffTarifsI {
	return s.staffTarifs
}
