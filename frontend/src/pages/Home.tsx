import { Link } from "react-router-dom"

const Home = () => {
  return (
    <div className="flex flex-col items-center justify-center h-[100vh] w-[100vw] bg-slate-900 text-white">
      <h1 className="text-4xl font-bold">Welcome to Ganja</h1>
      <p>Everything you need to create a server native application</p>
      <Link to={"/plan"} className="text-blue-600">See Plan</Link>
    </div>
  )
}

export default Home