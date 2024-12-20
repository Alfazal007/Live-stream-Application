import { UserContext } from '@/context/UserContext'
import { Suspense, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import StreamList from './StreamList'
import axios from "axios"
import { toast } from '@/hooks/use-toast'
import Navbar from './Navbar'

const Landing = () => {
	const navigate = useNavigate()
	const { user } = useContext(UserContext)
	const [streamList, setStreamList] = useState<{ id: string, creatorName: string }[]>([])

	async function fetchStreams() {
		const fetchStreamsUrl = 'http://localhost:8000/api/v1/stream/get-streams'
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
		console.log(response.data)
		setStreamList(response.data)
	}

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
		fetchStreams()
	}, [user])

	return (
		<>
			<Navbar />
			<div className="container mx-auto p-4">
				<h1 className="text-2xl font-bold mb-4">Available Streams</h1>
				<Suspense fallback={<div>Loading streams...</div>}>
					<StreamList streams={streamList} displayJoin={true} />
				</Suspense>
			</div>
		</>
	)
}

export default Landing
