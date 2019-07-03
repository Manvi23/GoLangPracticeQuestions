//implement singleton
packgage statemanager

type statemanager struct{
	state string
}

var singleton *manager //global pointer
var once sync.Once 

//wenver Getmanager gets called , singleton gets returned
//but a new initilization wont be called
//initializaton wld happen on;y once
func GetManager()*manager{
	once.Do(func(){singleton = &manager{state:"off"}}) //called anonymous function
	return singleton
}

func (sm *manager)Getstate()string{
	return sm.state
}

func (sm *manager)SetState(s string){
	sm.state = s
}
