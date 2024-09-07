"use client";

import { Recipe } from "@/types/recipe";
import {
  Button,
  Card,
  CardFooter,
  Image,
  Modal,
  ModalBody,
  ModalContent,
  ModalFooter,
  ModalHeader,
  useDisclosure,
} from "@nextui-org/react";

interface RecipeCardProps {
  recipe: Recipe;
}

export default function RecipeCard({ recipe }: RecipeCardProps) {
  const { isOpen, onOpen, onOpenChange } = useDisclosure();
  return (
    <div>
      <Card isPressable onPress={onOpen}>
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
      <Modal
        isOpen={isOpen}
        onOpenChange={onOpenChange}
        isDismissable={false}
        isKeyboardDismissDisabled={true}
      >
        <ModalContent>
          {(onClose) => (
            <>
              <ModalHeader className="flex flex-col gap-1">
                {recipe.name}
              </ModalHeader>
              <ModalBody>
                {recipe.meal}
                <p>Time to Cook</p>
                {recipe.cook_duration}
                <p>Instructions</p>
                {recipe.instructions}
              </ModalBody>
              <ModalFooter>
                <Button color="danger" variant="light" onPress={onClose}>
                  Close
                </Button>
                <Button color="primary" onPress={onClose}>
                  Action
                </Button>
              </ModalFooter>
            </>
          )}
        </ModalContent>
      </Modal>
    </div>
  );
}
