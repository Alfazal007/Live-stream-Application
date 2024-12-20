import CreateStreamButton from "./CreateStreamButton"
import Navbar from "./Navbar"

const StartStream = () => {
	return (
		<>
			<Navbar />
			<div className="mt-2">
				<CreateStreamButton />
			</div>
			<div>StartStream</div>
		</>
	)
}

export default StartStream
