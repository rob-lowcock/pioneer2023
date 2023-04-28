export default function Login() {
    return <div className="bg-bggray h-screen">
        <div className="px-4 sm:px-0 mx-auto container max-w-xl">
            <h1 className="font-bold text-2xl text-center pt-16 ">Sign in to Pioneer</h1>
            <form>
                <div>
                    <div className="block font-semibold text-sm mt-4 mb-2 text-black"><label htmlFor="email">Email</label></div>
                    <input type="email" className="form-input block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" id="email" placeholder="joe.bloggs@example.com" />
                </div>
                <div>
                    <div className="block font-semibold text-sm mt-4 mb-2 text-black"><label htmlFor="password">Password</label></div>
                    <input type="password" className="form-input block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" id="password" />
                </div>
                <div><button type="submit" className="rounded-md bg-webscale mt-8 w-full py-4 text-white hover:bg-webscale-lighter">Sign in →</button></div>
            </form>
        </div>
    </div>
}