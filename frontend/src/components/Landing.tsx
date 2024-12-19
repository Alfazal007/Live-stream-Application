import { UserContext } from '@/context/UserContext'
import { Suspense, useContext, useEffect, useState } from 'react'
import { useNavigate } from 'react-router-dom'
import CreateStreamButton from './CreateStreamButton'
import StreamList from './StreamList'

const Landing = () => {
	const navigate = useNavigate()
	const { user } = useContext(UserContext)
	const [streamList, setStreamList] = useState<{ id: string, creatorName: string }[]>([])

	useEffect(() => {
		if (!user) {
			navigate("/signin")
			return
		}
	}, [user])

	return (
		<div className="container mx-auto p-4">
			<h1 className="text-2xl font-bold mb-4">Available Streams</h1>
			<CreateStreamButton />
			<Suspense fallback={<div>Loading streams...</div>}>
				<StreamList streams={streamList} />
			</Suspense>
		</div>
	)
}

export default Landing
