import { Card, CardContent, CardFooter } from '@/components/ui/card'
import { Button } from '@/components/ui/button'
import { useNavigate } from 'react-router-dom'

type StreamCardProps = {
	stream: {
		id: string,
		creatorName: string
	},
	displayJoin: boolean
}

export default function StreamCard({ stream, displayJoin }: StreamCardProps) {
	const navigate = useNavigate()

	const handleJoinStream = () => {
		navigate(`/stream/${stream.id}`)
		return
	}

	return (
		<Card className="flex flex-col justify-between">
			<CardContent className="pt-6">
				<h3 className="text-lg font-semibold mb-2">{stream.creatorName}</h3>
				<p className="text-sm text-gray-500">Stream ID: {stream.id}</p>
			</CardContent>
			<CardFooter>
				{
					displayJoin &&
					<Button onClick={handleJoinStream} className="w-full">Join Stream</Button>
				}
			</CardFooter>
		</Card>
	)
}

