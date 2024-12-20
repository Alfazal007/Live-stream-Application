import { useContext, useEffect, useRef, useState } from 'react'
import { Button } from "@/components/ui/button"
import { Input } from "@/components/ui/input"
import { ArrowLeft } from 'lucide-react'
import { UserContext } from '@/context/UserContext'
import { useNavigate, useParams } from 'react-router-dom'

interface Messages {
	sender: string,
	message: string
}

interface BroadCastIncomingMessage {
	TypeOfMessage: string,
	Message: string,
	Sender: string
}

export default function VideoChatUser() {
	const messagesEndRef = useRef<HTMLDivElement>(null)
	const [messages, setMessages] = useState<Messages[]>([])
	useEffect(() => {
		messagesEndRef.current?.scrollIntoView({ behavior: 'smooth' })
	}, [messages])
	const [inputMessage, setInputMessage] = useState('')
	const { user } = useContext(UserContext)
	const WS_URL = "ws://localhost:8001/ws"
	const [socket, setSocket] = useState<WebSocket | null>(null)
	const { streamId } = useParams()
	const navigate = useNavigate()
	const joinMessage = {
		"typeOfMessage": "JOINUSER",
		"message":
			`{ "userId": "${user?.id}","roomId": "${streamId}","accessToken": "${user?.accessToken}"}`
	}
	// init type checker
	useEffect(() => {
		if (!user) {
			navigate("/")
		}
		const socketNew = new WebSocket(WS_URL)
		socketNew.onopen = () => {
			socketNew.send(JSON.stringify(joinMessage))
		};

		socketNew.onmessage = async (event: any) => {
			console.log("message in")
			const message: BroadCastIncomingMessage = JSON.parse(event.data)
			switch (message.TypeOfMessage) {
				case "NORMALMESSAGE":
					setMessages((messages) => [...messages, { sender: message.Sender || "Anonym", message: message.Message }])
					break
			}
		}

		socketNew.onclose = () => {
			navigate("/")
		};
		setSocket(socketNew)
	}, [])

	const handleSendMessage = (e: React.FormEvent) => {
		e.preventDefault()
		const messageToSend = {
			"typeOfMessage": "TEXTMESSAGE",
			"message":
				`{ "userId": "${user?.id}","roomId": "${streamId}","message": "${inputMessage.trim()}", "username":"${user?.username}"}`
		}
		if (inputMessage.trim()) {
			socket?.send(JSON.stringify(messageToSend))
			setMessages([...messages, { sender: user?.username || "", message: inputMessage.trim() }])
			setInputMessage('')
		}
	}

	const handleLeave = () => {
		console.log('Leaving video chat')
		socket?.close()
		return
	}

	return (
		<div className="flex flex-col">
			<Button
				variant="ghost"
				className="self-start mb-4 flex items-center"
				onClick={handleLeave}
			>
				<ArrowLeft className="mr-2 h-4 w-4" />
				Leave
			</Button>
			<div className="flex flex-col md:flex-row gap-4">
				<div className="w-full md:w-2/3">
					<video
						className="w-full aspect-video bg-gray-200"
						controls
						poster="/placeholder.svg?height=480&width=640"
					>
						<source src="/placeholder.mp4" type="video/mp4" />
						Your browser does not support the video tag.
					</video>
				</div>
				<div className="w-full md:w-1/3 flex flex-col">
					<div
						className="flex overflow-y-scroll bg-gray-100 p-4 rounded-lg mb-4"
						style={{ height: '600px' }}
					>
						<div className="flex flex-col">
							{messages.map(({ message, sender }, index) => (
								<div key={index} className="mb-2">
									<span className="font-bold">{sender}:</span> {message}
								</div>
							))}
							<div ref={messagesEndRef} />
						</div>
					</div>
					<form onSubmit={handleSendMessage} className="flex gap-2">
						<Input
							type="text"
							value={inputMessage}
							onChange={(e) => setInputMessage(e.target.value)}
							placeholder="Type your message..."
							className="flex-grow"
						/>
						<Button type="submit">Send</Button>
					</form>
				</div>
			</div>
		</div>
	)
}

