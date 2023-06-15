import { Bars2Icon, CheckIcon, PlusIcon } from "@heroicons/react/24/outline";
import Retrocolumn from "../retro/retrocolumn";
import Retrocard from "../retro/retrocard";
import { RetrocardType, createRetrocard, getLatestBoard } from "../services/retro";
import { redirect, useLoaderData, useRevalidator } from "react-router-dom";
import { UnauthorizedError } from "../utilities/errors";
import { useInterval } from "usehooks-ts";

export async function action({ request }: { request: Request }) {
    const formData = await request.formData();
    const title = formData.get("title");
    const col = formData.get("col");
    const errors = {
        message: "",
    }
    
    const retrocard: RetrocardType = {
        id: null,
        title: title?.toString() ?? "",
        column: parseInt(col?.toString() ?? "0"),
        active: true,
    }

    try {
        await createRetrocard(retrocard);
    } catch (e : any) {
        errors.message = e.message;
    }

    return {errors};
}

export async function loader() {
    try {
        const retrocards = await getLatestBoard();
    } catch (e : any) {
        if (e instanceof UnauthorizedError) {
            return redirect("/login");
        }
        console.log(e);
    }
    const cards = await getLatestBoard();
    const columns = [
        {
            id: 0,
            name: "I'm happy about...",
            cards: cards.data.filter((card: any) => card.attributes.column === 0),
        },
        {
            id: 1,
            name: "I'm wondering about...",
            cards: cards.data.filter((card: any) => card.attributes.column === 1),
        },
        {
            id: 2,
            name: "I'm sad about...",
            cards: cards.data.filter((card: any) => card.attributes.column === 2),
        },
    ];

    return { columns };
}

export default function Retro() {
    const { columns } = useLoaderData();
    const revalidator = useRevalidator();

    useInterval(() => {
        if (revalidator.state === "idle") {
            revalidator.revalidate();
        }
    }, 1000);

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
        { columns.length ? (
        <div className="md:columns-3 md:gap-0">
            { columns.map((column: any) => (
                <Retrocolumn column={column.id} title={column.name} key={column.id}>
                    { column.cards.length ? (
                        <>
                            { column.cards.map((card: any) => (
                                <Retrocard content={card.attributes.title} key={card.id} />    
                            ))}
                        </>
                    ) : (
                        <>
                            <p className="text-center text-sm text-gray-500">Nothing here yet! Be the first to add something!</p>
                        </>
                    )}
                </Retrocolumn>
            ))}
        </div>
        ) : (
            <></>
        )}
    </div>
}