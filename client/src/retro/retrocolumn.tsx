import { PropsWithChildren } from "react";
import RetroAddButton from "./retroaddbutton";

type RetroColumnProps = {
    title: string
    column: number
}

export default function Retrocolumn(props: PropsWithChildren<RetroColumnProps>) {
    return <div className="break-before-column p-4 m-4 bg-gray-50 rounded-lg">
        <div className="text-gray-500 font-title font-bold mb-4">
            <h2 className="font-title font-bold">{props.title} <span className="bg-gray-200 text-xs rounded-full ml-1 w-4 h-4 inline-block text-center">{props.column + 1}</span></h2>
        </div>
        <RetroAddButton col={props.column} />
        {props.children}
    </div>
}