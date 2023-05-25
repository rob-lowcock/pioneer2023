import { PropsWithoutRef } from "react"

type RetroCardProps = {
    content: string
}

export default function Retrocard(props: PropsWithoutRef<RetroCardProps>) {
    return <div className="p-6 mb-4 bg-white text-sm rounded-md border border-gray-200">
        {props.content}
    </div>
}