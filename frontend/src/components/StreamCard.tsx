import { Card, CardContent, CardFooter } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { Stream } from 'stream'
import { useNavigate } from 'react-router-dom'

type StreamCardProps = {
	stream: {
		id: string,
		creatorName: string
	}
}

export default function StreamCard({ stream }: StreamCardProps) {
	const navigate = useNavigate()

	const handleJoinStream = () => {
		// In a real application, you would join the stream here
		// and then redirect to the stream page
		navigate(`/stream/${stream.id}`)
	}

	return (
		<Card className="flex flex-col justify-between">
			<CardContent className="pt-6">
				<h3 className="text-lg font-semibold mb-2">{stream.creatorName}</h3>
				<p className="text-sm text-gray-500">Stream ID: {stream.id}</p>
			</CardContent>
			<CardFooter>
				<Button onClick={handleJoinStream} className="w-full">Join Stream</Button>
			</CardFooter>
		</Card>
	)
}

