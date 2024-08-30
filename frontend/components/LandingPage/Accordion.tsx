import React from 'react'
import {
    Accordion,
    AccordionContent,
    AccordionItem,
    AccordionTrigger,
} from "@/components/ui/accordion"


const AccordionFaq = () => {
    const faq = [
        {
            question: "How do I get started with Switchr?",
            answer: "Sign up for a free account, create your first project, and start managing your feature flags right away."
        },
        {
            question: "What is a feature flag?",
            answer: "A feature flag is a mechanism to turn functionality on or off during runtime, allowing for flexible feature management without redeployment."
        },
        {
            question: "How does Switchr ensure real-time updates?",
            answer: "Switchr leverages Redis to store and update feature flags instantly, ensuring that any changes you make are immediately applied."
        },
        {
            question: "Is there a limit to the number of projects or feature flags I can create?",
            answer: "Currently there is no limit to create flags, but it may change in future."
        }
    ];

    return (
        <Accordion type="single" collapsible>
            {faq.map((x, i) => {
                return (
                    <AccordionItem key={i} value={`item-${i}`}>
                        <AccordionTrigger className='text-xl text-left'>{x.question}</AccordionTrigger>
                        <AccordionContent className='text-xl text-primary/70'>
                            {x.answer}
                        </AccordionContent>
                    </AccordionItem>
                )
            })}
        </Accordion>

    )
}

export default AccordionFaq