package net.chrigel.clustercode.statemachine.actions;

import net.chrigel.clustercode.event.RxEventBus;
import net.chrigel.clustercode.statemachine.Action;
import net.chrigel.clustercode.statemachine.StateContext;
import net.chrigel.clustercode.statemachine.states.State;
import net.chrigel.clustercode.statemachine.states.StateEvent;
import net.chrigel.clustercode.transcode.TranscodingService;
import net.chrigel.clustercode.transcode.messages.TranscodeFinishedEvent;
import net.chrigel.clustercode.transcode.TranscodeTask;

import javax.inject.Inject;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.ExecutionException;

public class TranscodeAction extends Action {

    private final TranscodingService service;
    private CompletableFuture<TranscodeFinishedEvent> future;

    @Inject
    TranscodeAction(RxEventBus eventBus, TranscodingService service) {
        this.service = service;

        eventBus.register(TranscodeFinishedEvent.class, this::onTranscodeFinished);
    }

    private void onTranscodeFinished(TranscodeFinishedEvent event) {
        future.complete(event);
    }

    @Override
    protected StateEvent doExecute(State from, State to, StateEvent event, StateContext context) {
        log.entry(from, to, event, context);

        future = new CompletableFuture<>();

        service.transcode(TranscodeTask
            .builder()
            .media(context.getSelectedMedia())
            .profile(context.getSelectedProfile())
            .build());

        try {
            TranscodeFinishedEvent result = future.get();
            context.setTranscodeFinishedEvent(result);
            return result.isCancelled() ? StateEvent.CANCELLED : StateEvent.FINISHED;
        } catch (InterruptedException | ExecutionException e) {
            return StateEvent.ERROR;
        }
    }

}
