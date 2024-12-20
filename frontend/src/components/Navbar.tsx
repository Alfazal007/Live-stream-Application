import { useContext, useState } from 'react'
import { Menu, X } from 'lucide-react'
import { Button } from "@/components/ui/button"
import { Link } from 'react-router-dom'
import { UserContext } from '@/context/UserContext'

export default function Navbar() {
	const [isMenuOpen, setIsMenuOpen] = useState(false)

	const toggleMenu = () => {
		setIsMenuOpen(!isMenuOpen)
	}
	const { setUser } = useContext(UserContext)

	function onLogout() {
		setUser(null)

	}

	return (
		<nav className="bg-background shadow-md">
			<div className="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
				<div className="flex items-center justify-between h-16">
					<div className="flex items-center">
						<Link to="/" className="flex-shrink-0">
							<span className="text-2xl font-bold text-primary">StreamApp</span>
						</Link>
					</div>
					<div className="hidden md:block">
						<div className="ml-10 flex items-baseline space-x-4">
							<Button asChild variant="ghost">
								<Link to="/stream">Streams</Link>
							</Button>
							<Button variant="destructive" onClick={onLogout}>Logout</Button>
						</div>
					</div>
					<div className="md:hidden">
						<button
							onClick={toggleMenu}
							className="inline-flex items-center justify-center p-2 rounded-md text-primary hover:text-primary-foreground hover:bg-primary focus:outline-none focus:ring-2 focus:ring-inset focus:ring-white"
						>
							<span className="sr-only">Open main menu</span>
							{isMenuOpen ? (
								<X className="block h-6 w-6" aria-hidden="true" />
							) : (
								<Menu className="block h-6 w-6" aria-hidden="true" />
							)}
						</button>
					</div>
				</div>
			</div>

			{isMenuOpen && (
				<div className="md:hidden">
					<div className="px-2 pt-2 pb-3 space-y-1 sm:px-3">
						<Button asChild variant="ghost" className="w-full justify-start">
							<Link to="/stream">Streams</Link>
						</Button>
						<Button variant="destructive" className="w-full justify-start" onClick={onLogout}>Logout</Button>
					</div>
				</div>
			)}
		</nav>
	)
}

