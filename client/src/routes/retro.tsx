import { Bars2Icon, CheckIcon, PlusIcon } from "@heroicons/react/24/outline";
import Retrocard from "../retro/retrocard";
import Retrocolumn from "../retro/retrocolumn";

export default function Retro() {
    return <div className="h-screen bg-white">
        <nav className="p-6 border-b border-brdgray text-webscale">
            <div className="flex justify-between">
                <h1 className="text-xl font-bold font-title text-center text-webscale px-4 lowercase">Retro</h1>
                <button className="px-3 mr-4 text-black"><Bars2Icon className="h-6 w-6" /></button>
            </div>
        </nav>
        <div className="p-4 mt-4 mx-4 bg-webscale text-white rounded-lg">
            <h2 className="font-bold text-sm">We're talking about:</h2>
            <p className="text-2xl font-title text-center p-2 pb-4">Are we doing testing the right way?</p>
            <div className="flex flex-row">
                <a href="#" className="block basis-1/2 text-sm mr-2 text-center bg-webscale-lighter hover:bg-webscale-lighter2 p-2 rounded-lg"><PlusIcon className="h-6 w-6 inline-block" /> Add an action</a>
                <a href="#" className="block basis-1/2 text-sm ml-2 text-center bg-webscale-lighter hover:bg-webscale-lighter2 p-2 rounded-lg"><CheckIcon className="h-6 w-6 inline-block" /> Mark as done</a>
            </div>
        </div>
        <div className="md:columns-3 md:gap-0">
            <Retrocolumn column={0} title="I'm happy about...">
                <Retrocard content="A sample card" />
                <Retrocard content="We launched the product! 🎉" />
            </Retrocolumn>
            <Retrocolumn column={1} title="I'm wondering about...">
                <Retrocard content="Are we doing testing the right way?" />
                <Retrocard content="What does the restructure mean?" />
            </Retrocolumn>
            <Retrocolumn column={2} title="I'm sad about...">
                <Retrocard content="Redundancies" />
            </Retrocolumn>
        </div>
    </div>
}