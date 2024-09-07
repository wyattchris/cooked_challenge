"use client";

import { Recipe } from "@/types/recipe";
import { Card, CardFooter, Image } from "@nextui-org/react";

interface RecipeCardProps {
  recipe: Recipe;
}

export default function RecipeCard({ recipe }: RecipeCardProps) {
  return (
    <Card>
      <Image
        alt="Woman listing to music"
        className="object-cover"
        height={200}
        src={recipe.image_url}
        width={200}
      />
      <CardFooter>
        <div className="flex justify-between items-center">
          {recipe.name} {recipe.cook_duration}
        </div>
      </CardFooter>
    </Card>
  );
}
