import { FaceSmileIcon } from "@heroicons/react/24/outline";
import { PlusIcon, XMarkIcon } from "@heroicons/react/24/solid";
import { useRef, useState } from "react";
import Picker from '@emoji-mart/react';
import { Form, useActionData } from "react-router-dom";

export default function RetroAddButton(props: any) {
    const [showForm, setShowForm] = useState(false);
    const [showPicker, setShowPicker] = useState(false);
    const inputRef = useRef<HTMLInputElement>(null);
    const action = useActionData();

    const toggleForm = (e : React.MouseEvent<HTMLAnchorElement>) => {
        setShowForm(!showForm);
        setShowPicker(false);
    }

    const togglePicker = (e : React.MouseEvent<HTMLAnchorElement>) => {
        setShowPicker(!showPicker);
    }

    const pickerData = async () => {
        const response = await fetch('https://cdn.jsdelivr.net/npm/@emoji-mart/data');
        return response.json();
    }

    const pickerEvent = (emoji: any) => {
        inputRef.current!.value += emoji.native
    }

    if (action?.errors?.message == "" && inputRef.current != null) {
        inputRef.current!.value = "";
    }

    return <>
    { showForm ?
        <Form method="post" className="bg-gray-200 rounded-lg p-2 mb-2">
            <a href="#" onClick={toggleForm} className="float-right text-sm mb-1 text-center text-gray-500 bg-gray-200 hover:bg-gray-100 py-1 px-2 rounded"><XMarkIcon className="h-3 w-3 inline-block" /></a>
            <p className="p-1 text-gray-500 text-sm">What's on your mind?</p>
            <input type="text" ref={inputRef} className="block w-full rounded-md p-4 border-brdgray focus:border-slate focus:ring-0" autoFocus name="title" />
            <input type="hidden" name="col" value={props.col} />
            <a href="#" title="Toggle emoji picker" className="float-left text-gray-500 mt-2" onClick={togglePicker}><FaceSmileIcon className="h-6 w-6 inline-block" /></a>
            <button type="submit" className="block rounded-md bg-gray-500 ml-auto px-4 py-2 mt-2 text-sm text-white hover:bg-webscale-lighter">Add</button>
        </Form>
        :
        <a href="#" className="block text-sm mb-2 text-center text-gray-500 bg-gray-200 hover:bg-gray-100 p-2 rounded-lg" onClick={toggleForm}><PlusIcon className="h-6 w-6 inline-block" /> Add something</a>
    }
    { showPicker ?
        <div className="absolute">
        <Picker data={pickerData} onEmojiSelect={pickerEvent} />
        </div>
        : null
    }
    </>
}