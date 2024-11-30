import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Button} from "@/components/ui/button";
import React from "react";
import {Input} from "@/components/ui/input";
import {PlusIcon} from "lucide-react";
import {NextPage} from "next";
import {ShoppingList} from "@/types";
import {CreateList} from "@/components/create-list";

const Home: NextPage = async () => {
    const response = await fetch(`${process.env.NEXT_PUBLIC_API_URL}/api/shopping-lists`, {
        cache: 'no-store', // Чтобы данные всегда были актуальными
    });

    const shoppingLists: ShoppingList[] = await response.json();


    return (
        <div className="container mx-auto px-4 py-8">
            <h1 className="text-2xl font-bold mb-6">Мои списки покупок</h1>
            <div className="space-y-4">
                {shoppingLists.map((list) => (
                    <Card key={list.id} className="w-full">
                        <CardHeader>
                            <CardTitle>{list.name}</CardTitle>
                            <CardDescription>{list.items.length} items</CardDescription>
                        </CardHeader>
                        <CardFooter className="flex justify-between">
                            <Button variant="outline">Редактировать</Button>
                            <Button>Открыть</Button>
                        </CardFooter>
                    </Card>
                ))}
                <CreateList />
            </div>
        </div>
    );
}
export default Home