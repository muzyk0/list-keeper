import { NextResponse } from 'next/server';
import { PrismaClient } from '@prisma/client';
import { ShoppingList } from '@/types';

const prisma = new PrismaClient();

export async function GET() {
    try {
        const shoppingLists = await prisma.shoppingList.findMany({
            include: { items: true },
            orderBy: {
                id: 'desc'
            },
            take: 10
        });
        return NextResponse.json<ShoppingList[]>(shoppingLists);
    } catch {
        return NextResponse.json({ error: 'Failed to fetch shopping lists' }, { status: 500 });
    }
}


export async function POST(request: Request) {
    try {
        const body = await request.json();
        const newList = await prisma.shoppingList.create({

            data: {
                name: body.name,
                items: {
                    create: body.items,
                },
            },
        include: { items: true },
    });
        return NextResponse.json<ShoppingList>(newList, { status: 201 });
    } catch {
        return NextResponse.json({ error: 'Failed to create shopping list' }, { status: 500 });
    }
}

