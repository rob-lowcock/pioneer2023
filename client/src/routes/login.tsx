import { Form, redirect, useActionData } from "react-router-dom";
import { attemptLogin } from "../services/auth";
import { XCircleIcon } from "@heroicons/react/24/outline";

export async function action({ request }: { request: Request }) {
    const formData = await request.formData();
    const email = formData.get("email");
    const password = formData.get("password");
    const errors = {
        message: "",
    }

    if (!email || !password) {
        errors.message = "Email and password are required";
        // Flag something is wrong
        return errors;
    }
    try {
        const token = await attemptLogin(email!.toString(), password!.toString());
        console.log(token.data.id)
    } catch (error: any) {
        errors.message = error.message;
    }

    return errors;
}

export default function Login() {
    const error = useActionData() as any;

    return <div className="bg-bggray h-screen">
        <div className="px-4 sm:px-0 mx-auto container max-w-xl">
            { error?.message && <div className="pt-8 "><p className="p-4 bg-red text-white rounded-md"><XCircleIcon className="align-top inline h-6 w-6 mr-2 text-white" />{ error?.message }</p></div> }
            <h1 className="font-bold text-2xl text-center pt-16 ">Sign in to Pioneer</h1>
            <Form method="post" id="login-form">
                <div>
                    <div className="block font-semibold text-sm mt-4 mb-2 text-black"><label htmlFor="email">Email</label></div>
                    <input type="text" className="form-input block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" id="email" name="email" placeholder="joe.bloggs@example.com" autoComplete="username" />
                </div>
                <div>
                    <div className="block font-semibold text-sm mt-4 mb-2 text-black"><label htmlFor="password">Password</label></div>
                    <input type="password" className="form-input block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" id="password" name="password" autoComplete="current-password" />
                </div>
                <div><button type="submit" className="rounded-md bg-webscale mt-8 w-full py-4 text-white hover:bg-webscale-lighter">Sign in →</button></div>
            </Form>
        </div>
    </div>
}