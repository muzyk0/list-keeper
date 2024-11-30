import {NextResponse} from "next/server";
import {PrismaClient} from "@prisma/client";

const prisma = new PrismaClient();

export async function DELETE(request: Request, {params}: {params: {id: string}}) {
    try {
        const { id }: { id: string } = params;
        await prisma.shoppingList.delete({
            where: { id: Number(id) },
        });
        return NextResponse.json({ message: 'Shopping list deleted successfully' });
    } catch {
        return NextResponse.json({ error: 'Failed to delete shopping list' }, { status: 500 });
    }
}