package pkg

func fn() {
	for {
		switch {
		case true:
			break // MATCH /ineffective break statement/
		}
	}

	for {
		select {
		case <-ch:
			break // MATCH /ineffective break statement/
		}
	}

	for {
		switch {
		case true:
		}

		switch {
		case true:
			break // MATCH /ineffective break statement/
		}

		switch {
		case true:
		}
	}

	for {
		switch {
		case true:
			if true {
				break // MATCH /ineffective break statement/
			} else {
				break // MATCH /ineffective break statement/
			}
		}
	}

	for {
		switch {
		case true:
			if true {
				break
			}

			println("do work")
		}
	}

label:
	for {
		switch {
		case true:
			break label
		}
	}
}
