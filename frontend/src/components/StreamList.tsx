import StreamCard from "./StreamCard"

export default function StreamList({ streams }: { streams: { creatorName: string, id: string }[] }) {
	return (
		<div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-4">
			{streams.map((stream) => (
				<StreamCard key={stream.id} stream={stream} />
			))}
		</div>
	)
}

