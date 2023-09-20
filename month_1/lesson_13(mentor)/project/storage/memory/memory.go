package memory

import "backend_bootcamp_17_07_2023/lesson_8/project/storage"

type store struct {
	branches *branchRepo
	staffes  *staffRepo
}

func NewStorage() storage.StorageI {
	return &store{
		branches: NewBranchRepo(),
		staffes:  NewStaffRepo(),
	}
}

func (s *store) Branch() storage.BranchesI {
	return s.branches
}

func (s *store) Staff() storage.StaffesI {
	return s.staffes
}
