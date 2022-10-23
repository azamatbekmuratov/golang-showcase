// timout in channel
package main

func main(){
	for alive := true; alive; {
		timer := time.NewTimer(time.Hour)
		select {
		case news := 'some case':
			timer.Stop()
			fmt.Println(news)
		case <-timer.C:
			alive = false
			fmt.Println("No news in an hour. Service aborting.")
		}
	}
}
