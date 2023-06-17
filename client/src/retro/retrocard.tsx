import { PropsWithoutRef } from "react"
import { RetrocardType } from "../types/retrocard"

type RetroCardProps = {
    content: RetrocardType
    selectHandler: (newCard: RetrocardType) => void
}

export default function Retrocard(props: PropsWithoutRef<RetroCardProps>) {
    return <div className="p-6 mb-4 bg-white text-sm rounded-md border border-gray-200" onClick={() => props.selectHandler(props.content)}>
        {props.content.title}
    </div>
}