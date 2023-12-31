package service

type MachineService struct {
}

func NewMachineService() *MachineService {
	return &MachineService{}
}

func (s *MachineService) GetMachineInfo() error {
	// cpuCount, err := cpu.Counts(true)
	// if err != nil {
	// 	slog.Error(fmt.Sprintf("get cpu count error: %s", err.Error()))
	// 	return err
	// }
	// fmt.Println(cpuCount)
	return nil
}
