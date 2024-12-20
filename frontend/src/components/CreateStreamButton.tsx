import { Button } from '@/components/ui/button'
import { useNavigate } from 'react-router-dom'
import axios from "axios";
import { useContext } from 'react';
import { UserContext } from '@/context/UserContext';
import { toast } from '@/hooks/use-toast';

export default function CreateStreamButton() {
	const navigate = useNavigate()
	const { user } = useContext(UserContext)

	const handleCreateStream = async () => {
		const newStreamUrl = "http://localhost:8000/api/v1/stream/create-stream"
		const response = await axios.post(newStreamUrl, {}, {
			headers: {
				Authorization: `Bearer ${user?.accessToken}`
			}
		})
		if (response.status != 200) {
			toast({ title: "Issue creating the stream" })
			return
		}
		const startStreamUrl = "http://localhost:8000/api/v1/stream/start-stream"
		const startStreamResponse = await axios.put(startStreamUrl, {
			streamId: response.data.ID
		}, {
			headers: {
				Authorization: `Bearer ${user?.accessToken}`
			}
		})
		if (startStreamResponse.status != 200) {
			toast({ title: "Issue starting the stream" })
			return
		}
		navigate(`/admin/stream/${response.data.ID}`)
	}

	return (
		<Button onClick={handleCreateStream} className="mb-4">
			Create a Stream
		</Button>
	)
}

