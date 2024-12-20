import { Suspense, useContext, useEffect, useState } from "react"
import CreateStreamButton from "./CreateStreamButton"
import Navbar from "./Navbar"
import { UserContext } from "@/context/UserContext"
import { useNavigate } from "react-router-dom"
import StreamList from "./StreamList"
import { toast } from "@/hooks/use-toast"
import axios from "axios"

const StartStream = () => {
	const { user } = useContext(UserContext)
	const [streamList, setStreamList] = useState<{ creatorName: string, id: string }[]>([])
	const navigate = useNavigate()
	useEffect(() => {
		if (!user) {
			navigate("/")
			return
		}
		getUserStream()
	}, [])

	async function getUserStream() {
		const fetchStreamsUrl = 'http://localhost:8000/api/v1/stream/get-my-streams'
		const response = await axios.get(fetchStreamsUrl, {
			headers: {
				Authorization: `Bearer ${user?.accessToken}`
			}
		})
		if (response.status != 200) {
			toast({
				title: "Issue fetching the streams data"
			})
			return
		}
		setStreamList(response.data)
	}

	return (
		<>
			<Navbar />
			<div className="container mx-auto p-4">
				<div className="mt-2">
					<CreateStreamButton />
				</div>
				<h1 className="text-2xl font-bold mb-4">Available Streams</h1>
				<Suspense fallback={<div>Loading streams...</div>}>
					<StreamList streams={streamList} displayJoin={false} />
				</Suspense>
			</div>
		</>
	)
}

export default StartStream
