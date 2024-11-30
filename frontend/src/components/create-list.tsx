'use client'

import {Card, CardContent, CardHeader, CardTitle} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import React from "react";
import {Input} from "@/components/ui/input";
import {PlusIcon} from "lucide-react";
import {NextPage} from "next";

export const CreateList: NextPage = () => {
    const [newListName, setNewListName] = React.useState('');

    const createShoppingList = async () => {
        const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/shopping-lists`, {
            method: 'POST',
            body: JSON.stringify({name: newListName, items: [{name: 'item1'}, {name: 'item2'}]}),
            cache: 'no-store', // Чтобы данные всегда были актуальными
        });
    }

    return (
        <Card className="w-full">
            <CardHeader>
                <CardTitle>Создать новый список</CardTitle>
            </CardHeader>
            <CardContent>
                <div className="flex flex-col space-y-4">
                    <Input
                        placeholder="Название списка"
                        value={newListName}
                        onChange={(e) => setNewListName(e.target.value)}
                    />
                    <Button className="w-full"
                            onClick={createShoppingList}
                    >
                        <PlusIcon className="mr-2 h-4 w-4"/> Создать список
                    </Button>
                </div>
            </CardContent>
        </Card>
    );
}