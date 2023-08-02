package helpers

func (p Participant) GetParticipant(key int) {
	var participant = []Participant{
		{name: "Mufida", address: "Bandung", job: "Backend developer", reason: "belajar golang"},
		{name: "Nuha", address: "Garut", job: "Anak pertama", reason: "unlocked new skill"},
		{name: "Salimah", address: "Buah Batu", job: "Staff", reason: "mengisi waktu"},
	}

	// if(participant[key] == null){
		// notFound()
	// } else {
		printBiodata(participant[key], key)
	// }
}
