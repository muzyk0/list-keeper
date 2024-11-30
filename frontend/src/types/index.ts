export interface ShoppingItem {
    id: number;
    name: string;
    completed: boolean;
    shoppingListId: number;
}

export interface ShoppingList {
    id: number;
    name: string;
    items: ShoppingItem[];
    createdAt: Date;
    updatedAt: Date;
}