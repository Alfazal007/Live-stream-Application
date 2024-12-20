import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { SignUp } from "./components/Signup";
import { SignIn } from "./components/Signin";
import UserProvider from "./context/UserContext";
import Landing from "./components/Landing";
import './index.css'
import StartStream from "./components/StartStream";
import VideoChatUser from "./components/UserWatchingStream";
import VideoChatAdmin from "./components/AdminStreamScreen";

export interface User {
	accessToken: string;
	username: string;
	id: string;
}

function App() {
	const router = createBrowserRouter([
		{
			path: "/signup",
			element: <SignUp />,
		},
		{
			path: "/signin",
			element: <SignIn />,
		},
		{
			path: "/",
			element: <Landing />,
		},
		{
			path: "/stream",
			element: <StartStream />,
		},
		{
			path: "/stream/:streamId",
			element: <VideoChatUser />
		},
		{
			path: "/admin/stream/:streamId",
			element: <VideoChatAdmin />
		}
	]);

	return (
		<>
			<UserProvider>
				<RouterProvider router={router} />
			</UserProvider>
		</>
	);
}

export default App;
