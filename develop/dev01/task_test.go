package main_test

import "time"

func main() {
	ch := make(chan int, 2)

	ch <- 42

	// заполненная структура hchan
	//ch = {chan int}
	//qcount = {uint} 1 - количество элем в буфере
	//dataqsiz = {uint} 2 - размер буфера
	//*buf = {*[2]int} len:2 - определяет буфер с данными, кольцевой буфер
	//elemsize = {uint16} 8 - размер элемента в канале
	//closed = {uint32} 0 - канал в данный момент открыт. 1- закрыт
	//*elemtype = {*runtime._type} - указатель на тип данных
	//sendx = {uint} 1 - индексы смещения для записи
	//recvx = {uint} 0 - индекс смещения для чтения
	//recvq = {waitq<int>} - односвязный список очереди из горутин на чтение
	//sendq = {waitq<int>} - односвязный список очереди из горутин на запись
	//lock = {runtime.mutex} - мьютекс используемый для операций изменения состояния канала

	nonbufCh := make(chan bool)
	go func() {
		time.Sleep(time.Second)
		nonbufCh <- true
	}()
	<-nonbufCh
}
