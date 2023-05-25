import { PlusIcon, XMarkIcon } from "@heroicons/react/24/solid";
import { useState } from "react";

export default function RetroAddButton() {
    const [showForm, setShowForm] = useState(false);

    const onClick = (e : React.MouseEvent<HTMLAnchorElement>) => {
        setShowForm(!showForm);
    }

    return <>
    { showForm ?
        <form action="#" className="bg-gray-200 rounded-lg p-2 mb-2">
            <a href="#" onClick={onClick} className="float-right text-sm mb-1 text-center text-gray-500 bg-gray-200 hover:bg-gray-100 py-1 px-2 rounded"><XMarkIcon className="h-3 w-3 inline-block" /></a>
            <p className="p-1 text-gray-500 text-sm">What's on your mind?</p>
            <input type="text" className="block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" autoFocus />
            <button type="submit" className="block rounded-md bg-gray-500 ml-auto px-4 py-2 mt-2 text-sm text-white hover:bg-webscale-lighter">Add</button>
        </form>
        :
        <a href="#" className="block text-sm mb-2 text-center text-gray-500 bg-gray-200 hover:bg-gray-100 p-2 rounded-lg" onClick={onClick}><PlusIcon className="h-6 w-6 inline-block" /> Add something</a>
    }
    </>
}