
/*
// subscribing for APPROVAL events. We use watchers without fast-forwarding past events
func SubscribeForApprovals(session *passport.PassportSession, listenChannel chan<- *passport.PassportPassportApproved) (event.Subscription, error) {
	subscription, err := session.Contract.WatchPassportApproved(&bind.WatchOpts{
		Start:   nil, //last block
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}




// 
func SubscribeForRecent(session *jschan.session, listenChannel chan<- *jschan.GetManageRecentResponse, board_name string) (event.Subscription, error) {
	subscription, err := session.jschan.WatchRecent(&bind.WatchOpts{
		Board:   board_name, 
		Context: nil, // nil = no timeout
	}, listenChannel,
	)
	if err != nil {
		return nil, err
	}
	return subscription, err
}


						subscription, err := SubscribeForRecent(session, ch, board_name) // 
						if err != nil {
							log.Println(err)
						}

	//go AsyncApproveChain(ctx, subscription, update.Message.From.ID, auth, passportCenter, session, userDatabase)

EventLoop:
	for {
		select {
		case <-ctx.Done():
			{
				subscription.Unsubscribe()
				break EventLoop
			}
		case eventResult := <-ch:
			{
				fmt.Println("Recent post message: ", eventResult.Message)
				



				subscription.Unsubscribe()
				break EventLoop
			}
		}
	}





*/