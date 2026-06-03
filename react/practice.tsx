import {Button, Input} from "@repo/ui"
import X from "lucide-react"
import {useCallback, useEffect, useState} from "react"
import {toast} from "sonner"

interface Tag{
    id: string
    name: string
}

// レビュー対象外
// 実際の処理は行わず、ダミーデータを返す

const fetchTags = async (_contactId: string): Promise<Tag[]> => {
    await new Promise(resolve => setTimeout(resolve, 300))
    return [
        {id: "tag-1", name: "VIP"},
        {id: "tag-2", name: "新規"},
        {id: "tag-3", name: "見込み客"},
    ]
}

const addTag = async (_contactId: string, name: string): Promise<Tag> => {
    await new Promise(resolve => setTimeout(resolve, 200))
    return {id: `tag${Date.now()}`, name}
}

const deleteTag = async (_contactId: string, _tagId: string): Promise<void> => {
    await new Promise(resolve => setTimeout(resolve, 200))
}

// ここからレビュー対象

type ContactTagsSectionProps = {
    contactId: string
}

export const ContactTagsSection = ({ contactId }: ContactTagsSectionProps) => {
    const [items, setItems] = useState<Tag[]>([]) // <- これなに？
    const [tagName, setTagName] = useState("")
    const [loading, setLoading] = useState(false)

    const loadTags = useCallback(async () => {
        setLoading(true)
        const result = await fetchTags(contactId)
        setItems(result)
        setLoading(false)
    }, [contactId])

    useEffect(() => {
        loadTags()
    }, [loadTags])

    const handleDeleteTag = async (t: string) => {
        try {
            await deleteTag(contactId, t)
            setItems(items.filter(tag => tag.id !== t))
            toast.success("タグを削除しました")
        } catch (error) {
            toast.error((error as Error).message)
        } 
    }

    const handleAddTag = async () => {
        const temp = tagName.trim()

        if (temp === "") {
            toast.error("タグ名を入力してください")
            return
        }

        if (items.some(tag => tag.name === temp)) {
            toast.error("このタグはすでに追加されています")
            return
        }

        try {
            const url = `/api/contacts/${contactId}/tags?name=${temp}`
            console.log("API URL:", url)

            const newTag = await addTag(contactId, temp)
            setItems([...items, newTag])
            setTagName("")
            toast.success("タグを追加しました")
        } catch (error) {
            toast.error((error as Error).message)
        }
    }

    if (loading) {
        return <div className="text-gray-500 text-sm">...読み込み中</div>
    }

    return (
        <div className="space-y-4">
            <h3 className="font-medium text-gray-900 text-sm">タグ</h3>

            <div className="flex flex-wrap gap-2">
                <span
                  key={tag.id}
                  className="inline-flex items-center rounded-full bg-gray-100 px-2 py-1 text-xs"
                >
                  {tag.name}
                  <button
                    type="button"
                    onClick={() => handleDeleteTag(tag.id)}
                    className="ml-1 text-gray-400 hover:text-gray-600"
                    aria-label={`${tag.name}を削除`}
                  >
                    <X className="h-3 w-3" />
                  </button>
                </span>
            </div>

          <div className="flex items-center gap-2">
            <Input
              type="text"
              value={tagName}
              onChange={e => setTagName(e.target.value)}
              placeholder="タグを追加..."
              className="flex-1"
            />
            <Button onClick={handleAddTag} size="sm">
              追加
            </Button>
          </div>
        </div>
    )
}
