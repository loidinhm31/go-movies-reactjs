import Stripe from "stripe";
import { withoutRole } from "src/libs/auth";

const handler = withoutRole("banned", async (req, res) => {
    try {
        const payment = req.body;

        const stripePromise = new Stripe(`${process.env.STRIPE_PRIVATE_KEY}`, {
            apiVersion: "2022-11-15",
        });

        const paymentIntent = await stripePromise.paymentIntents.create(payment);

        // Send publishable key and PaymentIntent details to client
        res.send({
            clientSecret: paymentIntent.client_secret,
            paymentId: paymentIntent.id,
        });
    } catch (e) {
        return res.status(400).send({
            message: e.message,
        });
    }
});

export default handler;
