import { Button } from '@/components/ui/button'
import { useNavigate } from 'react-router-dom'

export default function CreateStreamButton() {
	const navigate = useNavigate()

	const handleCreateStream = () => {
		// TODO:: send api request before redirect
		navigate('/new-stream')
	}

	return (
		<Button onClick={handleCreateStream} className="mb-4">
			Create a Stream
		</Button>
	)
}

